# Hajimari 🌅

Hajimari is a sleek and intuitive start page for your Kubernetes cluster. It
dynamically discovers and organizes applications, supports personalized
configurations, and integrates with Kubernetes resources for a streamlined
interface.

![Hajimari](assets/screen01.png)

![User Configuration](assets/screen02.png)

![App Groups](assets/screen03.png)

## Features

- **Search Functionality**: Web and application search with customizable search
  providers.
- **Dynamic Application** Discovery: Automatically lists applications from
  Kubernetes Ingresses or Custom Resources.
- **Replica Status Monitoring**: Displays real-time replica status for ingress
  endpoints.
- **Custom App Support**: Add non-Kubernetes applications via configuration.
- **Bookmark Management**: Organize and customize bookmark lists.
- **Theming Options**: Choose from selectable themes or create custom themes
  with automatic light/dark mode support.
- **User Customization**: Per-user or per-browser configuration overrides.
- **Multi-Instance Support**: Run multiple Hajimari instances with targeted
  application displays.

## Installation

### Helm

1. Add the Hajimari Helm repository:

```bash
helm repo add hajimari https://hajimari.io
```

2. Update the Helm repository:

```bash
helm repo update
```

3. Install Hajimari:

```bash
helm install hajimari hajimari/hajimari
```

For detailed Helm chart configuration, refer to the
[Helm documentation](charts/hajimari).

### Kustomize

Refer to the [k8s](./k8s) folder for resources needed for Kustomize.

### Locally

To build Hajimari locally, clone the repository and run:

```bash
make deps
make build
```

Or for local development with hot reload:

```bash
make deps
make dev
```

**Requirements**

- Go 1.19
- Node.js 18

> [!NOTE]
>
> Hajimari requires access to a kubeconfig file for a service account with
> permissions to access `Ingress` and `EndpointSlice` objects.

## Configuration

### Kubernetes Ingresses

Hajimari discovers applications by reading specific annotations on Kubernetes
Ingresses. Add the following annotations to your Ingress resources:

| Annotation                | Description                                                                | Required |
| ------------------------- | -------------------------------------------------------------------------- | -------- |
| `hajimari.io/enable`      | Set to `true` to display the application in Hajimari.                      | Yes      |
| `hajimari.io/icon`        | Icon name from Material Design Icons.                                      | No       |
| `hajimari.io/appName`     | Custom application name (overrides Ingress name).                          | No       |
| `hajimari.io/group`       | Custom group name (overrides namespace-based grouping).                    | No       |
| `hajimari.io/instance`    | Comma-separated list of Hajimari instance names for multi-instance setups. | No       |
| `hajimari.io/url`         | Custom URL for the application (must include `http://` or `https://`).     | No       |
| `hajimari.io/targetBlank` | Set to `true` to open links in a new tab/window.                           | No       |
| `hajimari.io/info`        | Short description of the application.                                      | No       |

### Configuration Options

Hajimari supports configuration via `ConfigMap` or `values.yaml` (for Helm).
Below are the available options:

| Field                   | Description                                             | Default      | Type                |
| ----------------------- | ------------------------------------------------------- | ------------ | ------------------- |
| `instanceName`          | Name of the Hajimari instance.                          | `""`         | `String`            |
| `defaultEnable`         | Expose all Ingresses in selected namespaces by default. | `false`      | `Bool`              |
| `namespaceSelector`     | Select namespaces using names or label selectors.       | `any: true`  | `NamespaceSelector` |
| `name`                  | Name used in the greeting.                              | `"You"`      | `String`            |
| `title`                 | Browser title for the instance.                         | `"Hajimari"` | `String`            |
| `lightTheme`            | Theme for light mode.                                   | `"gazette"`  | `String`            |
| `darkTheme`             | Theme for dark mode.                                    | `"horizon"`  | `String`            |
| `customThemes`          | List of custom themes.                                  | `[]`         | `[]Theme`           |
| `showGreeting`          | Toggle greeting and date display.                       | `true`       | `Bool`              |
| `showAppGroups`         | Group applications by namespace.                        | `true`       | `Bool`              |
| `showAppUrls`           | Display application URLs.                               | `true`       | `Bool`              |
| `showAppInfo`           | Show application descriptions.                          | `false`      | `Bool`              |
| `showAppStatus`         | Display replica status indicators.                      | `true`       | `Bool`              |
| `showBookmarkGroups`    | Group bookmarks by category.                            | `true`       | `Bool`              |
| `showGlobalBookmarks`   | Show global bookmarks on custom start pages.            | `false`      | `Bool`              |
| `alwaysTargetBlank`     | Open apps/bookmarks in a new tab by default.            | `false`      | `Bool`              |
| `defaultSearchProvider` | Default search provider name.                           | `"Google"`   | `String`            |
| `searchProviders`       | List of custom search providers.                        | See below    | `[]SearchProvider`  |
| `customApps`            | List of custom applications.                            | `[]`         | `[]AppGroup`        |
| `globalBookmarks`       | List of bookmark groups.                                | `[]`         | `[]BookmarkGroup`   |

### Custom Resources (HajimariApp)

For non-Ingress setups (e.g., Istio Virtual Services or Traefik
`IngressRoutes`), define applications using Kubernetes Custom Resources:

```yaml
apiVersion: hajimari.io/v1alpha1
kind: Application
metadata:
  name: hajimari-issues
spec:
  name: Hajimari Issues
  group: info
  icon: simple-icons:github
  url: https://github.com/mahyarmirrashed/hajimari/issues
  info: Submit issues for this project
  targetBlank: true
```

### Namespace Selector

The `namespaceSelector` field controls which namespaces Hajimari monitors:

| Field           | Description                                                  | Default | Type                   |
| --------------- | ------------------------------------------------------------ | ------- | ---------------------- |
| `any`           | Select all namespaces if `true`.                             | `false` | `Bool`                 |
| `labelSelector` | Filter namespaces using a Kubernetes `metav1.LabelSelector`. | `null`  | `metav1.LabelSelector` |
| `matchNames`    | List of specific namespace names.                            | `null`  | `[]String`             |

> [!NOTE]
>
> Combining `labelSelector` and `matchNames` results in a union of matched
> namespaces.

### Custom Themes

Define custom themes with the following attributes:

| Field             | Description             | Type     |
| ----------------- | ----------------------- | -------- |
| `name`            | Theme name.             | `String` |
| `backgroundColor` | Background color (hex). | `String` |
| `primaryColor`    | Primary color (hex).    | `String` |
| `accentColor`     | Accent color (hex).     | `String` |

### Search Providers

Customize search providers to override defaults:

| Field       | Description                                                 | Type     |
| ----------- | ----------------------------------------------------------- | -------- |
| `name`      | Search provider name.                                       | `String` |
| `token`     | Token to activate the provider in the search bar.           | `String` |
| `icon`      | Icon name or URL for the provider.                          | `String` |
| `searchUrl` | Search URL with `{query}` placeholder for the query string. | `String` |
| `url`       | URL to redirect to when only the token is entered.          | `String` |

#### Default Search Providers

```yaml
searchProviders:
  - name: Google
    token: g
    icon: simple-icons:google
    searchUrl: https://www.google.com/search?q={query}
    url: https://www.google.com
  - name: DuckDuckGo
    token: d
    icon: simple-icons:duckduckgo
    searchUrl: https://duckduckgo.com/?q={query}
    url: https://duckduckgo.com
  - name: IMDB
    token: i
    icon: simple-icons:imdb
    searchUrl: https://www.imdb.com/find?q={query}
    url: https://www.imdb.com
  - name: Reddit
    token: r
    icon: simple-icons:reddit
    searchUrl: https://www.reddit.com/search?q={query}
    url: https://www.reddit.com
  - name: YouTube
    token: y
    icon: simple-icons:youtube
    searchUrl: https://www.youtube.com/results?search_query={query}
    url: https://www.youtube.com
  - name: Spotify
    token: s
    icon: simple-icons:spotify
    searchUrl: https://open.spotify.com/search/{query}
    url: https://open.spotify.com
```

### Custom Apps

Add external or non-Ingress applications via the `customApps` configuration:

| Field   | Description                   | Type     |
| ------- | ----------------------------- | -------- |
| `group` | Group name (e.g., namespace). | `String` |
| `apps`  | List of custom applications.  | `[]App`  |

#### App

| Field         | Description               | Type     |
| ------------- | ------------------------- | -------- |
| `name`        | Application name.         | `String` |
| `icon`        | Icon name or URL.         | `String` |
| `url`         | Application URL.          | `String` |
| `info`        | Short description.        | `String` |
| `targetBlank` | Open in a new tab/window. | `Bool`   |

### Bookmark Groups

Organize bookmarks with the following structure:

| Field       | Description          | Type         |
| ----------- | -------------------- | ------------ |
| `group`     | Bookmark group name. | `String`     |
| `bookmarks` | List of bookmarks.   | `[]Bookmark` |

#### Bookmark

| Field         | Description               | Type     |
| ------------- | ------------------------- | -------- |
| `name`        | Bookmark name.            | `String` |
| `icon`        | Icon name or URL.         | `String` |
| `url`         | Bookmark URL.             | `String` |
| `targetBlank` | Open in a new tab/window. | `Bool`   |

## Usage

### Search Functionality

The search bar supports:

- **Application Filtering**: Use `/` to filter applications in real-time.
- **Custom Search Providers**: Start queries with `@<token>` to trigger a
  provider’s `searchUrl` with the query replacing `{query}`.
- **Quick Access**: Enter only `@<token>` to redirect to the provider’s `url`.

### Icons

Use Iconify icons (e.g., `mdi:kubernetes`, `simple-icons:google`) or provide a
direct image URL for any `icon` field.

### Custom Start page Setup

1. Open Hajimari in your browser and click the hamburger menu (bottom left).
2. Edit settings in the built-in YAML editor.
3. Save to generate a custom start page with a unique URL ID.
4. Set this URL as your browser’s homepage or new tab page. Recommended
   extensions:
   - Firefox: New Tab Override
   - Chrome: Custom New Tab URL

> [!NOTE]
>
> Hajimari does not include authentication. Consider restricting access via an
> `Ingress` controller.

> [!NOTE]
>
> Hajimari does not auto-update when configuration map is changed. It is
> recommended to use the [Reload Operator](https://github.com/stakater/Reloader)
> to automatically restart the Hajimari deployment when the configuration map
> changes.

## Contributing

We welcome contributions! For significant changes, please open an issue to
discuss your proposal. Ensure tests are updated as needed. Run `make help` for
details on linting, testing, and other tasks.

## About

Hajimari (始まり), meaning "beginnings" in Japanese, is designed as a browser
start page, making it the perfect entry point for your Kubernetes cluster.

### Acknowledgments

- [SUI](https://github.com/jeroenpardon/sui) For the great start page template.
- [Forecastle](https://github.com/stakater/Forecastle) Ideas for integrating K8s
  ingress.

## License

[Apache-2.0](https://choosealicense.com/licenses/apache-2.0/)
