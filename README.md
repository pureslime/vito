# 💊 VITO
**The extensible, pill-based CLI engine.** _VITO_ is a platform designed to centralize all your tooling with a "Pills" (plugin) system which are lightweight and powerfull, executed on a native way (Go) or through WebAssembly (Any other language).

## 🚀 Main Features
- **Zero-Config**: It auto-initializes itself on first execution.
- **Pill Engine**: Automatic discovery of toolins in `~/.vito/pills`.
- **Dual Runner**: It supports native binaries (Go) and modules of **WebAssembly** (Wasm).
- **Built-in SDK**: It helps you create your own toolings with a consistent UI and auto-documentation.
- **Upgradable**: You can upgrade VITO directly from the core just by using `vito upgrade`.

## 📦 Quick Installation
Copy and paste this command in your terminal to install _VITO_.

**Linux / MacOS**
```bash
curl -sSL https://raw.githubusercontent.com/pureslime/vito/main/install.sh | sh
```

**Windows**
_For now not supported_.

That will install the binary on `~/.vito/bin` and add the route to your PATH automatically.

## 🛠️ How it works
_VITO_ looks for files in you folder called: `pills`. By executing any command, _VITO_ registers them dynamically:
```bash
# Look on installed pills
vito --help
# Execute a command from a specific pill (example)
vito net ping 8.8.8.8
```

## 🏗️ Create your own Pill (SDK)
Developing Pills on _VITO_ is extremely easy thanks to the SDK. Here you have an example of a network "Pill":
```go
package main
import (
	"github.com/pureslime/vito-sdk/go/sdk"
)

func main() {
	// Initializes the pill (main command) as "net"
	p := sdk.NewPill("net")
	
	// Will define a subcommand with its description and what the subcommand does.
	p.Handle("ping", "Checks connectivity", func(args []string) error) {
		// Your Logic will be Here
		return nil
	}
	
	// Will put the pill automatically on the help and for you to be used on VITO automatically.
	// It autogenerates a manifest that VITO uses.
	p.Run()
}
```

## 📂 File Structure
_VITO_ maintains all organized in one place:
- `~/.vito/bin`: VITO's binary code.
- `~/.vito/config`: Configuration files on YAML.
- `~/.vito/pills`: You toolings (binaries or `.wasm`)

## ✨ Contributions
Refer to [CONTRIBUTIONS.md]("./CONTRIBUTIONS.md") to know more on how to contribute to _VITO_.

## 🗒️ License
Distributed under CC BY-NC-SA 4.0 + DISCLAIMERS. Look at [LICENSE](./LICENSE) for more information.

Made with ❤️ by [pureslime](https://github.com/pureslime)
