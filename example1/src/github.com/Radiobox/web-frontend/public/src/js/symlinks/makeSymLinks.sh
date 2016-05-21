rm -dr ./*.js 
ln -s ../../../vendor/bower/lib/requirejs-domready/domReady.js ./domReady.js
ln -s ../../../vendor/bower/lib/angular/angular.js ./angular.js
ln -s ../../../vendor/bower/lib/angular-ui-router/release/angular-ui-router.js ./angular-ui-router.js
ln -s ../../../vendor/bower/lib/angular-bootstrap/ui-bootstrap-tpls.js ./ui-bootstrap-tpls.js
ln -s ../../../vendor/bower/lib/requirejs/require.js ./require.js
ln -s ../../../vendor/bower/lib/angular-cookies/angular-cookies.js ./angular-cookies.js
ln -s ../../../vendor/bower/lib/angular-animate/angular-animate.js ./angular-animate.js
ln -s ../../../vendor/bower/lib/ng-file-upload/angular-file-upload.js angular-file-upload.js
ln -s ../../../vendor/bower/lib/ng-file-upload/angular-file-upload-shim.js angular-file-upload-shim.js
ln -s ../../../vendor/bower/lib/angular-sanitize/angular-sanitize.js angular-sanitize.js
ln -s ../../../vendor/bower/lib/moment/moment.js moment.js
ln -s ../../../vendor/bower/lib/moment-timezone/moment-timezone.js moment-timezone.js
ln -s ../../../vendor/bower/lib/moment-timezone/moment-timezone.json moment-timezone.json


#TODO, Fix this - issues with Facebook's MIME type causing all kinds of hell, so we fetch from them manually
curl -o facebook-all.js http://connect.facebook.com/en_US/all.js

#The JWPlayer Manually downloaded
curl -o jwplayer.js http://jwpsrv.com/library/1dQE4kslEeOG+RIxOQfUww.js
