Radiobox Web Frontend js-app
===============
#### Version 0.1 ####

Radiobox Frontend Angular JS App

## Testing ##

Work in Progress. 

Grunt will be used to test builds

Karma will be used for automated tests


## Node NPM##

We try and keep all external dependencies in Node modules should be installed in the `vendor/node_modules` foler. Make sure and add `--prefix ./vendor/node_modules` at the end of  `npm install`.

## The Angular/Require App ##

### Bower Package Management ###

Bower configuration is set in `app/bower.json`. Note that ui-router may not be compatible with the bleeding edge version of angular.

### Require JS ###

Require loads initial libs installed by bower in the `app/lib` folder, which includes

*  angular: 1.2.1
*  requirejs-domready: 2.0.1
*  angular-ui-router: 0.2.0

The JS libs are fetched asynchronously, but executed in the order of dependence.

### Angular ###

`ng-app="app"` is not set in the HTML markup, but is instead attached to `window.document` in `app/js/bootstrap.js`.

Flow of Require.js:

    js/main.js -> js/bootstrap.js -> js/app.js -> Myriad of folders: [js/controller,js/directives, js/filters, js/services]
