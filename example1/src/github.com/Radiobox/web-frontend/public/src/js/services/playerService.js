define(['./module', 'jwplayer'], function (services, jwplayer) {
    'use strict';

    services.provider('playerService', function () {


        this.$get = function($rootScope, $timeout, $sce) {

            var id = 'mainPlayer';

            $rootScope.$watch(function () { return wrappedService.volume }, function (newVal, oldVal) {
                if (typeof newVal !== 'undefined' && newVal !== oldVal) {
                    jwplayer(id).setVolume(newVal);
                }
            });
            var mainPlayerOptions = {
//                width: 1,
//                height: 1,
                playlist: [
                        {
                            file: "/media/blank.mp3",
                            artist: "Radiobox",
                            title: "Choose a Track"
                        }
                ],
                primary: "flash",
                autostart: false
            }

            jwplayer(id).setup(mainPlayerOptions);


            //Event Linkers
            jwplayer(id).onReady(function () {
                wrappedService.setVolume(jwplayer(id).getVolume());
                wrappedService.state = jwplayer(id).getState();
                wrappedService.position = "";
                wrappedService.duration = "";
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onTime(function (event) {
                //Execute only if on whole second
                if (event.position % 1 === 0 ) {
                    var duration = jwplayer(id).getDuration();

                    if (duration === -1){
                        wrappedService.duration = $sce.trustAsHtml('&infin;');
                    } else if (angular.isNumber(duration)) {
                        wrappedService.duration = $rootScope.seconds2HMS(duration);
                    }

                    wrappedService.position = $rootScope.seconds2HMS(event.position);
                }
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onPause(function () {
                wrappedService.state = 'PAUSED';
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onBuffer(function () {
                wrappedService.state = 'BUFFERING';
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onPlay(function () {
                wrappedService.state = 'PLAYING';

                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onPlaylist(function (playlist) {
                wrappedService.playlistQueued ={};
                wrappedService.playlistCurrent = playlist;
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onPlaylistItem(function (arg) {
                var nowplay = jwplayer(id).getPlaylistItem(arg.index);

                if(angular.isString(nowplay.title)) wrappedService.playing.title = nowplay.title;
                if(angular.isString(nowplay.artist)) wrappedService.playing.artist = nowplay.artist;
                if(angular.isString(nowplay.pic_square)) wrappedService.playing.pic_square = nowplay.pic_square;

                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onIdle(function () {
                wrappedService.state = 'IDLE';
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onPlaylistComplete(function () {
                if ($rootScope.isEmptyObj(wrappedService.playlistQueued)) {
                    jwplayer(id).load(wrappedService.playlistCurrent);
                }
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onComplete(function () {
                wrappedService.file = {};
                $timeout(function(){ $rootScope.$apply(); });
            });

            jwplayer(id).onBufferChange(function (buffer) {
                wrappedService.buffer = buffer;
                $timeout(function(){ $rootScope.$apply(); });
            });

            var wrappedService = {
                /**
                 * Public Methods
                 */

                setLoadFile: function(link, type) {
                    this.setLoadList({file:link});
                },

                setLoadList: function(array, autoplay, startOn) {
                    this.is_init = true;
                    this.playlistQueued = array;
                    jwplayer(id).load(this.playlistQueued);
                    if(startOn >= 0) wrappedService.setPlaylistItem(startOn);
                    if (autoplay) jwplayer(id).play(true);
                },

                setPlaylistItem: function(index) {
                    //if playing and request sent here does not match current, then load the playList
                    jwplayer(id).playlistItem(index);
                },

                play: function() {
                    jwplayer(id).play(true);
                },

                playToggle: function() {
                    jwplayer(id).play();
                },

                pause: function() {
                    jwplayer(id).pause();
                },

                getPlaylistIndex: function() {
                    return jwplayer(id).getPlaylistIndex();
                },

                getPlaylistId: function() {
                    return jwplayer(id).getPlaylistItem().playlistId;
                },

                getPlayingCode: function() {
                    return jwplayer(id).getPlaylistItem().playlistId + "-" + jwplayer(id).getPlaylistIndex();
                },

                getVolume: function() {
                    return jwplayer(id).getVolume()
                },

                setVolume: function(vol) {
                    jwplayer(id).setVolume(vol);
                    this.volume = vol;
                },

                next: function() {
                    jwplayer(id).playlistNext();
                },

                prev: function() {
                    jwplayer(id).playlistPrev();
                },



                /**
                 * Public Properties
                 */
                volume: null,
                is_init: false,
                playing: {},
                state: null,
                position: null,
                duration: null,
                isLive: null,
                buffer: null,
                playlistQueued: {},
                playlistCurrent: {}

            }

            $rootScope.player = wrappedService;

            return wrappedService
        }

    });

});