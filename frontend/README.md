# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```sh
# create a new project
pnpm dlx sv create my-app
```

To recreate this project with the same configuration:

```sh
# recreate this project
pnpm dlx sv@0.12.4 create --template minimal --types ts --add prettier eslint vitest="usages:unit,component" playwright tailwindcss="plugins:typography,forms" sveltekit-adapter="adapter:static" devtools-json mcp="ide:other+setup:local" --install pnpm frontend
```

## Developing

Once you've created a project and installed dependencies with `pnpm install`, start a development server:

```sh
vp dev

# or start the server and open the app in a new browser tab
vp dev --open
```

## Building

To create a production version of your app:

```sh
vp build
```

You can preview the production build with `vp preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.
