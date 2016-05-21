define(['./module',], function (controllers) {
    'use strict';

    controllers.controller('helloCtrl', ['$scope', '$modal', '$log', '$http', '$filter', '$animate', '$cookieStore', 'userService', function ($scope, $modal, $log, $http, $filter, $animate, $cookieStore, userService) {
        $animate.enabled(false);
        $scope.foo = "bar";
        $scope.liveNow = [{
            connected: 564,
            category: 'dj',
            followers: '10k',
            genre: 'ELEC',
            artist: 'Drunken Hearts',
            venue: 'Fox Theater',
            location: 'Boulder, Colorado'
        },
        {
            connected: '23k',
            category: 'dj',
            followers: '121k',
            genre: 'ELE',
            artist: 'MARS LIMO',
            venue: 'Belly Up',
            location: 'Aspen, CO'
        },
        {
            connected: 22,
            category: 'dj',
            followers: '342',
            genre: 'TRP',
            artist: 'GFMW',
            venue: 'The Heights',
            location: 'Cleveland, OH'
        }];
        $scope.liveNowInterval = -1;
        
        $scope.featuredArtists = [
            {
                artist: 1,
                link: '/drunkenhearts',
                image: '/img/temp/banner3.png',
                title: 'Drunken Hearts',
                caption: '',
                paragraph: ''
            },
            {
                artist: 1,
                link: '/tumbleweedwanderers',
                image: '/img/temp/banner2.png',
                title: 'Tumbleweed Wanderers',
                caption: 'Oakland California',
                paragraph: ''
            },
            {
                artist: 1,
                link: '/',
                image: '/img/temp/banner1.png',
                title: 'Yarmonygrass',
                caption: 'Live from Rancho Del Rio, Colorado',
                paragraph: ''            }
        ]
        
        $scope.featuredArtistsInterval = 10000;
        
        $scope.featuredVenues = [
            {
                artist: 1
            }
        ]
        
        $scope.featuredVenuesInterval = -1;
        
        $scope.upcoming = [{
            connected: 564,
            category: 'dj',
            followers: '10k',
            genre: 'ELEC',
            artist: 'The Drunken Hearts',
            venue: 'Fox Theatre',
            location: 'Boulder, Colorado USA'
        }]

        $scope.upcomingInterval = -1;
        
        $scope.betaArtist = {
            artist_name: '',
            email: ''
        };
        $scope.openSignup = userService.openSignup;
        $scope.betaSubmitting = false;
        $scope.betaSubmitted = false;
        $scope.signupBeta = function(){
            if($scope.betaArtist.artist_name == ''){
                alert('Tell us what artist you represent');
                return false;
            }
            if($scope.betaArtist.email == ''){
                alert('We will need your email to contact you!');
                return false;
            }
            $scope.betaSubmitting = true;
            $http.post('/api/beta-signup', $scope.betaArtist).success(function(dat){
                $scope.betaSubmitted = true;
            }).error(function(dat){
                $scope.betaSubmitting = false;
                try {
                    alert(dat.notification.input.email);
                } catch(err) {
                    alert('Something went wrong, try again or email us');
                }
                
            })
            return false;
            
        };

   }]);

});
