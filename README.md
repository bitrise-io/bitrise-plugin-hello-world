# Bitrise CLI Hello World Plugin

**Bitrise CLI Example Plugin.**

Prints Hello World!.

## How to use this Plugin

Can be run directly with the [bitrise CLI](https://github.com/bitrise-io/bitrise), requires version 1.3.0 or newer.

First install the plugin:
`bitrise plugin-install --source https://github.com/bitrise-tools/bitrise-plugin-hello-world/releases/download/0.9.1/bitrise-plugin-hello-world-$(uname -s)-$(uname -m) --name hello`

After that, you can use it:
`bitrise :hello`

## Guidelines for Plugin Development

### Plugin inputs

* Plugin becomes informations from bitrise through `BITRISE_PLUGINS_MESSAGE` environment as JSON string.

   *Currently bitrise informations contains bitrise version, in format:* `{"version":"CURRENT_BITRISE_VERSION"}`

* Also command line arguments passed to the plugin, which can be used to control Plugin behavior.

   *Example: this plugin respons to:* `bitrise :hello requirements`

### Plugin outputs

* Plugin should send informations to bitrise through os.Stdout as JSON string.

   *Example: send requirements to bitrise:* `[{"ToolName":"bitrise","MinVersion":"1.3.0","MaxVersion":"1.3.0"}]`

* Plugin outputs and errors should printed to the STDERR

### Plugin requirements

Through plugin install bitrise ask plugin for requirements (i.e.: calls 'bitrise :my_plugin requirements') and waits for answer in following JSON format:

`[{"ToolName":"TOOL_NAME_1","MinVersion":"SEMANTIC_MIN_VERSION_1","MaxVersion":"SEMANTIC_MAX_VERSION_1"}, {"ToolName":"TOOL_NAME_2","MinVersion":"SEMANTIC_MIN_VERSION_2","MaxVersion":"SEMANTIC_MAX_VERSION_2"}, ...]`.

Currently supporting tool requirements (ToolName):

* **bitrise**
* **stepman**
* **envman**

Notes:

* Any missing tool (in plugin requirement list) will be skipped in check (i.e.: no requirements for this tool).
* Any empty version (in plugin requirement list) will be skipped in check (i.e.: no requirements for this type of version).

Example:

If your plugin requires only bitrise with minimum version "1.0.0", your plugin requirement list looks like:

`[{"ToolName":"bitrise","MinVersion":"1.0.0","MaxVersion":""}]`.
