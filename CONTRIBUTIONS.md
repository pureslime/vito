# 🤝 Contributions
I like that you want to help out VITO to get widely known and powerful! It doesn't matter if you're a Go expert or if you just started your journey. I'm looking for developers to improve the **Core** and content creators to expand out **Pills and Templates**.

_VITO_ was born to stop the "copy-paste" cycle of boilerplates. We want tools that feel fast, look good, and are easy to extend.

## 📦 Tech Stack
> THE BOX THE BOX! — Homer J. Simpson.

So... what's in the box?
- **CLI Engine:** [Cobra](https://github.com/spf13/cobra)
- **TUI Framework:** [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **Styling**: [Lip Gloss](https://github.com/charmbracelet/lipgloss)

If you are adding UI components, please stick to Charm's ecosystem to keep VITO lightweight and visually consistent.

## 📜 Conventions
These are established rules that I and everyone else should follow. Use a convention for your commits and specify what you're doing:
- Use a convention for your commits, please, specify what you're doing: `feat(): added in a major change on the core`.
  - `fix()`: Something that has been broken has been fixed.
  - `feat()`: You added a new feature
  - `docs():` Changed readme, or any related documentation of the project.
  - `refactor()`: Code changes that neither fix a bug nor add a feature.

## 💊 How can you help?
- **Creating templates**, Fell like a specific stack is missing? Build it! I'll review it and consider adding it to the official VITO library.
- **Developing Pills**, Add new functions to the VITO Engine. Whether you're a Go pro or prefer other tools, we want your ideas.
- **Reporting Bugs**, if something explodes (let's hope not!), please open an _Issue_.
- **Refining the UI**, Help out to make the terminal brilliant again with better styling and UX.

## 🛠️ First Steps
- **Local development** To test your changes locally, you can run:
```bash
go run main.go create my-test-project
```
And that works for every command you create. For example: `go run main.go audit.`

Don't be shy! You're welcome to push your PRs, even if they get denied, just try your best!

## 🛠️ Workflow
1. Fork the project
2. Create your feature branch with `git checkout -b feature/new-implementation-name`.
3. Commit your changes `git commit -m 'feat(): added this new thing'`.
4. Push the branch `git push origin feature/new-implementation-name`.
5. Open a Pull Request

## 💡 Wishlist / Ideas
- A Pill for auto-generating Dockerfiles.
- Template for Rust/Wasm vitamins.
- A Command to analyze a project health (`vito audit`)

## 📜 Code of Conduct
_VITO_ is an open and healthy community. I expect every collaborator to be constructive and respectful. **Everyone** is free to collaborate, share their opinion, and contribute, regardless of their background or the minority group they belong to. You are welcome here.

Please keep discussions focused on VITO-related topics and maintain a professional environment.

## 🤖 AI on the Code
**No AI Slop**. I don't support low-effort AI-generated code. While I use AI myself sometimes, I dont' want "AI Slop" in a serious project aimed at helping the community.

- ✅ **The VITO Way**: Using AI _as a tool_ to debug, brainstorm, or explain complex logic.
- ❎ **The AI Guru Way**: Generating an entire PR with AI and submitting it without understanding or refining it or tested it.

_And you don't wan to be an "AI Guru" of X (Ex-Twitter), right?_
_Remember when an AI Guru expert on generating AI Images got humbled by an official ONE PIECE animator there? We don't want you to feel the same._

_VITO_ is about **quality**, not generic fillers. If a PR feels like a raw AI output without human touch, it will be closed. **Keep it authentic.**
