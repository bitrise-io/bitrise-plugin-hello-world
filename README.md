# Bitrise CLI Hello World Plugin

Bitrise CLI Example Plugin

## How to use this Plugin

Can be run directly with the [bitrise CLI](https://github.com/bitrise-io/bitrise).

First install the plugin:
`bitrise plugin-install --source https://github.com/bitrise-tools/bitrise-plugin-hello-world/releases/download/0.1.0/bitrise-plugin-hello-world --name hello`

After that, you can use it:
`bitrise :hello`

## Guidelines for Plugin Development
  * Plugin outputs and errors should printed to the STDERR
  * bitrise informations are passed to the Plugin as JSON string on STDIN
  * Plugin can pass informations to bitrise through STDOUT as JSON string
