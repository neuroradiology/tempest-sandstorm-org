@0x9498f3818bafa387;
# This schema describes the APIs available from Sandstorm's externally
# facing websockets (as opposed to those available to grains).
#
# If your sandstorm server is sandstorm.example.net and has HTTPS enabled,
# you can connect to this API at the url:
#
#   wss://sandstorm.example.net/_capnp-api
#
# The bootstrap interface will be of type ExternalApi

using Go = import "/go.capnp";
$Go.package("external");
$Go.import("sandstorm.org/go/tempest/capnp/external");

using Util = import "util.capnp";
using Collection = import "collection.capnp";
using Spk = import "package.capnp";
using Grain = import "grain.capnp";
using Identity = import "identity.capnp";

interface ExternalApi {
  # The bootstrap interface when connecting to the externally facing
  # websocket

  getSessions @0 () -> Sessions;
  # Return handles to sessions for the current user. Fields representing
  # roles the user does not have will be left null. If the user is not logged
  # in, this will throw an exception. logged-in status is determined by
  # the HTTP headers used in the websocket connection request (A session
  # cookie for the web UI. TODO: we may want to support bearer tokens
  # or something for programmatic access?

  restore @1 (sturdyRef :Data) -> (cap :Capability);
  # Restore a sturdyRef as a live capability.

  authenticator @2 () -> (authenticator :Authenticator);
  # Return an authenticator that can be used to log in.
}

struct Sessions {
  visitor @0 :VisitorSession;
  user @1 :UserSession;
}

interface Authenticator {
  # An authenticator provides functionality for authenticating a user with
  # the Tempest server.

  sendEmailAuthToken @0 (address :Text);
  # Send an email authentication token to the specified email address.
  # Making an http request to /login/email/<base64url-encoded token> will
  # return a response that sets a login cookie. In the future, it will
  # be possible to also redeem the token via ExternalApi.restore(), returning
  # some appropriate capability.
}

interface VisitorSession {
  # A VisitorSession provides operations that only require the 'visitor' role.

  views @0 () -> (views :UiView.Keyring);
  # Get the ui views into grains that the caller has access to. The keys are
  # opaque base64-url encoded identifiers. These are stable for a given user,
  # but references to the same view held by different users may not have the
  # same key.
  #
  # In the case where the view is the root UiView of a grain, the key is the
  # same as the grain id.
}

struct Package {
  manifest @0 :Spk.Manifest;
  controller @1 :Controller;

  interface Controller {
    create @0 (title :Text, actionIndex :UInt32) -> (id :Text, view :UiView);
    # Create a new grain using this package, with the given title
    # and using the action at the specified index in manifest.Actions
    # to spawn. Right now the action must have input == none.
    # Returns the id of the new grain and it's primary UiView.
  }

  interface InstallStream extends (Util.ByteStream) {
    # An InstallStream is a Util.ByteStream into which an .spk file can
    # be written; when done() is called, the package will be installed
    # and become available to the user.

    getPackage @0 () -> (id :Text, package :Package);
    # getPackage returns the package once it is installed. Generally,
    # users will want to call this immediately, and then write the
    # spk using write. This will not return until after done() is called.
  }
}

interface UserSession {
  # A UserSession provides operations that require the 'user' role.

  installPackage @0 () -> (stream :Package.InstallStream);
  # Install a package. If the stream is dropped before calling done(), installation
  # is cancelled.

  listPackages @1 (into :Collection.Pusher(Text, Package));
  # List the packages that the caller has installed.
}

struct UiView {
  # A UiView includes information about and access to a Grain.UiView. For now,
  # this maps 1-to-1 onto grains, but in the future Tempest will support
  # interfacing with any implementation of the UiView interface, which will
  # enable for example exporting multiple UiViews from the same grain.

  title @1 :Text;
  # The title of the view

  sessionToken @2 :Text;
  # A session token which can be exchanged for a cookie on a ui-* subdomain. To
  # open a UI session in the browser, navigate to:
  #
  #   http(s)://ui-${subdomain}.sandstorm.example.net/_sandstorm-init?sandstorm-sid=${sessionToken}&path=${path}
  #
  # This will set an authentication cookie using the Set-Cookie header, and
  # return a redirect to /${path}
  #
  # The token is valid until `handle` is dropped.

  subdomain @3 :Text;
  # The the subdomain, sans ui- prefix, that a session for this grain view
  # should be loaded from.

  viewInfo @4 :Grain.UiView.ViewInfo;
  # View info for the UiView.

  controller @0 :Controller;
  # Controller for manipulating the grain. When controller is dropped, sessionToken
  # is invalidated.

  interface Controller {
    makeSharingToken @0 (permissions :Identity.PermissionSet, note :Text) -> (token :Text);
    # Make a sharing token, which can be used to construct sharing URLs. The url should be:
    # http(s)://sandstorm.example.net/#/shared/${token}
    #
    # The permissions argument specifies which permissions should be shared. Each index
    # corresponds to a permission from the permissions field of the ViewInfo.
    #
    # Note is a human-readable note regarding the token; this may be displayed in various
    # places in the UI.
    #
    # In addition to accessing the UiView from the UI, you can also pass it to
    # ExternalApi.restore, which will return a Util.Getter(Util.KeyValue(Text, UiView)),
    # where the key is as in the collection returned by VisitorSession.view().
  }

  interface Keyring extends (Collection.Puller(Text, UiView)) {
    # A Keyring represents the set of capabilities (always UiViews) which
    # belong directly to the user (as opposed to e.g. ones that are held by
    # grains). These are displayed in the user's grain list in the UI.

    attach @0 (controller :Controller);
    # Add a UiView to the keyring.
  }
}
