/**
 * loads sub modules and wraps them up into the main module
 * this should be used for top-level module definitions only
 */
define([
    'angular',
    'uiRouter',
    'uiBootstrap',
    'angular-cookies',
    'angular-animate',
    'angular-sanitize',
    './controllers/index',
    './directives/index',
    './filters/index',
    './services/index',
    'file-upload',
    'moment',
    'moment-timezone',
    'moment-timezone-data'
], function (ng) {
    'use strict';
    return ng.module(['app'], [
        'app.services',
        'app.controllers',
        'app.filters',
        'app.directives',
        'ui.router',
        'ui.bootstrap',
        'ngCookies',
        'ngAnimate',
        'ngSanitize',
        'angularFileUpload'
    ]);
});