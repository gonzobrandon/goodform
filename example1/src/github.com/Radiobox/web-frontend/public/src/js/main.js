/**
 * configure RequireJS
 * prefer named modules to long paths, especially for version mgt
 * or 3rd party libraries
 */
require.config({

    paths: {
        'domReady': 'symlinks/domReady',
        'angular': 'symlinks/angular',
        "uiRouter": "symlinks/angular-ui-router",
        "uiBootstrap": "symlinks/ui-bootstrap-tpls",
        "angular-cookies": "symlinks/angular-cookies",
        "angular-animate": "symlinks/angular-animate",
        "angular-sanitize": "symlinks/angular-sanitize",
        "file-upload": "symlinks/angular-file-upload",
        "file-upload-shim": "symlinks/angular-file-upload-shim",
        "facebook": "symlinks/facebook-all",
        "jwplayer": "symlinks/jwplayer",
        "moment": "symlinks/moment",
        "moment-timezone": "symlinks/moment-timezone",
        "moment-timezone-data": "symlinks/moment-timezone-data"
    },

    /**
     * for libs that either do not support AMD out of the box, or
     * require some fine tuning to dependency mgt'
     */
    shim: {
        'angular': {
            exports: 'angular',
            deps: ['file-upload-shim']
        },
        'uiRouter':{
            deps: ['angular']
        },
        'uiBootstrap':{
            deps: ['angular']
        },
        'angular-cookies':{
            deps: ['angular']
        },
        'file-upload':{
            deps: ['angular']
        },
        'angular-animate':{
            deps: ['angular']
        },
        'angular-sanitize':{
            deps: ['angular']
        },
        'facebook' : {
            exports: 'FB'
        },
        'jwplayer' : {
            exports: 'jwplayer'
        },
        'moment-timezone': {
        	deps: ['moment']
        },
        'moment-timezone-data': {
        	deps: ['moment-timezone']
        }
    },
    
    deps: [
        'bootstrap'
    ]
});
