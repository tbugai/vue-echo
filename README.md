# Vue.js + GO

Basic starting point for apps that use a Vue.js frontend with a Go backend ([Echo](https://github.com/labstack/echo) in this case).

## Usage

Clone the repository into a directory of your choosing on your local development machine.  The _Makefile_ will use the directory name as the name of the server executable generated.

The Makefile will watch for changes to the backend code and rebuild/launch it when necessary.  It also launches `yarn serve` for the frontend which will do hot reloading for the vue.js frontend.

```
make dev
```


## Distribution

When you are ready to ship your website, the **Makefile** will build the frontend for production and boundle it along with the backend into the `dist/` subfolder.
 
```
make dist
```

Additionally, you can get the contents of the `dist/` folder in a timestamped zip file

```
make zip
```
