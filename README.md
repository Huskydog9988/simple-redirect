# Simple Redirect

Permanently redirects users from the current URI to the new one.

Not only does it redirect them, but attempts to preserve their desired page during the redirect.

## Environment Variables

| Variable      | Description                                                                                                                                                                     |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `REDIRECT_To` | **(Required)** Specifies hostname to redirect too. _(Should be no trailing `/`)_                                                                                                |
| `PORT`        | Changes the port the server runs on                                                                                                                                             |
| `PERMANENT`   | Whether the redirect should be [permanent](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308) or [not](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307) |
