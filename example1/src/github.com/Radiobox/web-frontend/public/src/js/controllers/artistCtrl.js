define(['./module'], function (controllers) {
    'use strict';
    
    /* DISPLAY PAGE ROOT */
    controllers.controller('artistCtrl', ['$scope', '$modal', '$log', '$http', '$filter', '$animate', '$cookieStore', 'userService', 'rootService', 'playerService', function ($scope, $modal, $log, $http, $filter, $animate, $cookieStore, userService, rootService, playerService) {

        var artist = rootService.slugTarget;
        var playlist = [];
        $scope.Date = function(arg){
            return new Date(arg);
        };
        $scope.months = ['JAN','FEB','MAR','APR','MAY','JUN','JUL','AUG','SEP','OCT','NOV','DEC'];    
        
        $scope.artist = false;
        $scope.albums = false;
        $scope.events = false;
        $scope.is_live_event_now = false;


         $scope.playTrack = function(index){
            playerService.setLoadList(playlist, false, index);
        };

        $scope.playLive = function() {

            if (playerService.getPlaylistIndex() !== 0) {
                $scope.playTrack(0);
            } else if (playerService.getPlaylistIndex() === 0 && playerService.state === 'PLAYING') {
                playerService.pause();
            };
        };

        
        $http.get('/api/artists/'+ rootService.slugTarget.id + '?joins={"type":"full","albums":{"type":"full", "tracks":{"type":"full"}},"events":{"type":"full"}}').success(function(data){
            $scope.artist = data.response;
            $scope.albums = data.response.albums.response;
            $scope.events = data.response.events.response;
            
            var trackIndex = 0;
            playlist = [];

            //Check if live event is in progress, if yes - push the live stream on top of playlist
            $scope.liveEvent = data.response.events.response[0];
            if ($scope.liveEvent.is_in_progress) {

                $scope.is_live_event_now = true;

                var entry = {
                    file: $scope.liveEvent.allocated_broadcasts[0].client_hls,
                    artist: $scope.artist.username,
                    title: $scope.liveEvent.title,
                    pic_square: $scope.artist.pic_square,
                    isLive: true
                };

                playlist.push(entry);
                trackIndex++;

            }


            angular.forEach($scope.albums, function(val, key){
                angular.forEach(val.tracks.response, function(track, tkey){
                    var entry = {
                        file: track['media-links'].mp3,
                        artist: track.artist.username,
                        title: track.title,
                        isLive: false
                    };
                    track.playlistIndex = trackIndex++;
                    playlist.push(entry);
                });
            });

            if ($scope.is_live_event_now) {
                playerService.setLoadList(playlist, true, 0);
            }


        });

   }]);
   
    controllers.controller('trickArtistCtrl', ['$scope', '$http', '$filter', 'userService', 'rootService', '$interval', '$sce', '$rootScope', 'playerService', '$state',  function ($scope, $http, $filter, userService, rootService, $interval, $sce, $rootScope, playerService, $state) {

        $scope.player = playerService;

        // Fake arrays for easier templating. Remove after api implementation
        $scope.bsArray = [{s:1},{s:1},{s:1},{s:1},{s:1},{s:1},{s:1}];
        $scope.shortArray = [{s:1},{s:1},{s:1}]
        $scope.Date = function(arg){
            return new Date(arg.split('.')[0].replace(' ','T'));
        };
        $scope.months = ['JAN','FEB','MAR','APR','MAY','JUN','JUL','AUG','SEP','OCT','NOV','DEC'];        
        $scope.artist = false;

        var request = {};

        request.joins = {
            events_live: {
                page_size: 10,
                joins: {
                    venue: {
                        format: 'compact',
                    }
                }
            },
            albums: {
                page_size: 3,
                joins: {
                    tracks: {}
                }
            }
        }
        request.joins = encodeURIComponent(JSON.stringify(request.joins));


        var buildArtistPlaylist = function(data) {
            // BEGIN BUILD ARTIST PLAYLIST FOR PLAYER
            var artistPlaylist = [];
            var albumIndex = 0;
            var trackIndex = 0;
            var playlistIndex = 0;
            var autoplay = false;
            $scope.playlistId = 'artist-' + data.id;

            //Check if live event is in progress, if yes - push the live stream on top of playlist
            if (data.event_live.is_in_progress) {
                var tempTrack = {
                    file: data.event_live.event_live_provisioned_broadcast.client_hls,
                    artist: data.event_live.title,
                    title: data.event_live.title_from,
                    pic_square: data.event_live.pic_square,
                    isLive: true,
                    playlistId: $scope.playlistId
                };
                artistPlaylist.push(tempTrack);
                playlistIndex += 1;
                autoplay = true;
            }

            angular.forEach(data.albums, function(album) {
                trackIndex = 0;

                angular.forEach(album.tracks, function(track) {
                    var tempTrack = {
                        file: track.preview_media_url,
                        artist: album.artist_name,
                        title: track.title,
                        pic_square: album.pic_square,
                        playlistId: $scope.playlistId
                    };

                    data.albums[albumIndex].tracks[trackIndex].playlistIndex = playlistIndex;
                    playlistIndex += 1;
                    trackIndex += 1;
                    artistPlaylist.push(tempTrack);

                });

                albumIndex += 1;
            });

            return artistPlaylist;
        }

        $scope.playTrack = function(id) {
            if (playerService.getPlaylistId() !== $scope.playlistId || playerService.state !== 'PLAYING') {
                $scope.player.setLoadList($scope.artistPlaylist, true);
                $scope.player.setPlaylistItem(id);
            } else if (playerService.getPlaylistId() === $scope.playlistId) {
                $scope.player.setPlaylistItem(id);
            }

        };

        $scope.playLive = function() {
            if (playerService.getPlayingCode !== $scope.playlistId + "-0") {
                $scope.playTrack(0);
            } else if (playerService.getPlayingCode === $scope.playlistId + "-0" && playerService.state !== 'PLAYING') {
                playerService.pause();
            };
        };


        var getUrl = '/js/json/' + $state.current.name  + '.json?with='



        $http.get(getUrl).success(function(data){

            $scope.artistPlaylist = buildArtistPlaylist(data);

            // END BUILD ARTIST PLAYLIST FOR PLAYER
            $scope.artist = data;

        });

   }]);
   
   /* DISPLAY SUB PAGES */
   controllers.controller('artistFollowersCtrl', ['$scope', function ($scope) {
            
   }]);
   controllers.controller('artistAlbumsCtrl', ['$scope', function ($scope) {
            
   }]);
   controllers.controller('artistScheduleCtrl', ['$scope', function ($scope) {
            
   }]);
   
   /* CREATE ARTIST */
   
   controllers.controller('artistCreateCtrl', function($scope, $location) {
       $scope.onSubmit = function(dat){
           $location.url('/' + dat.response.slug);
       }
   });
   

});