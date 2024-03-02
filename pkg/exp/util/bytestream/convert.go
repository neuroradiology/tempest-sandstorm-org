package bytestream

// convert between util.ByteStream and io.WriteCloser.

import (
	"context"
	"io"

	"sandstorm.org/go/tempest/capnp/util"
)

// Convert an io.WriteCloser to a util.ByteStream. WriteCloser's Write and
// Close methods correspond to write and done, respectively. expectSize is
// a no-op.
func FromWriteCloser(wc io.WriteCloser) util.ByteStream {
	return util.ByteStream_ServerToClient(&wcByteStream{wc: wc})
}

// Implementation of util.ByteStream that wraps an io.WriteCloser; returned by
// FromWriteCloser().
type wcByteStream struct {
	wc io.WriteCloser
}

func (wc *wcByteStream) Write(ctx context.Context, p util.ByteStream_write) error {
	data, err := p.Args().Data()
	if err != nil {
		return err
	}
	_, err = wc.wc.Write(data)
	return err
}

func (wc *wcByteStream) Done(context.Context, util.ByteStream_done) error {
	return wc.wc.Close()
}

func (*wcByteStream) ExpectSize(context.Context, util.ByteStream_expectSize) error {
	// no-op; just here to satisfy the interface.
	return nil
}

// Convert a util.ByteStream to an io.WriteCloser. The WriteCloser's Write and
// Close methods correspond to the bytestream's write and done methods,
// respectively. 'ctx' will be used when making rpc calls on the bytestream.
//
// Note that the resulting object's Write method does not wait for the call to
// resolve, and thus cannot determine whether the write was successful. It will
// report an error if the context is canceled, and may report one if a previous
// call has resolved to an error.
func ToWriteCloser(ctx context.Context, stream util.ByteStream) io.WriteCloser {
	return &byteStreamWC{
		ctx:    ctx,
		stream: stream,
	}
}

// Implementation of io.WriteCloser that wraps a util.ByteStream; returned by
// ToWriteCloser().
type byteStreamWC struct {
	ctx    context.Context
	stream util.ByteStream
}

// Write calls Write on the ByteStream
func (bs byteStreamWC) Write(data []byte) (n int, err error) {
	err = bs.ctx.Err()
	if err != nil {
		return 0, err
	}
	err = bs.stream.Write(bs.ctx, func(p util.ByteStream_write_Params) error {
		p.SetData(data)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

// Close calls Done on the ByteStream
func (bs byteStreamWC) Close() error {
	fut, rel := bs.stream.Done(bs.ctx, func(util.ByteStream_done_Params) error {
		return nil
	})
	defer rel()
	_, err := fut.Struct()
	if err != nil {
		return err
	}
	return bs.stream.WaitStreaming()
}
