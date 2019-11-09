# Minify Your Web Assets without Node!

`minsrv` is a build tool that helps you minify and hash our CSS and Javascript files automatically, so that you don't have to set up webpack or any of the other Node tools.

It comes with two parts.

# The `minsrv` CLI

This is a CLI that you can use to compress your JS, CSS and HTML, and put everything into new, hashed files

>TODO: use [`pkger`](https://github.com/markbates/pkger) to embed everything into the binary

# Middleware

This is middleware that does a few things:

1. Intercepts HTML pages before they go back to the client
2. Parses all of the `<script>` and `<link>` tags & changes the `src` and `href` (respectively) URLs to point to the newly hashed files
3. Minifies the HTML and returns it to the client

