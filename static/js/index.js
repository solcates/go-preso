var app = angular.module('gopreso', [
    'ngMaterial',
    'ngAnimate',
    'ngAria',
    'ngMessages',
    'ngRoute',
    'gopreso.main.controller',
    'gopreso.login.controller',
])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/', {
            templateUrl: '/static/views/main.html',
            controller: 'MainController',
            controllerAs: "main",
            name: "main"
        });
    }])
    .controller('NavController', function ($scope, $http, $mdSidenav) {
        $scope.close = function () {
            // Component lookup should always be available since we are not using `ng-if`
            $mdSidenav('left').close()
                .then(function () {
                    $log.debug("close LEFT is done");
                });

        };


    })

    .run(['$rootScope', '$location', 'AuthService', function ($rootScope, $location, AuthService) {
        $rootScope.$on('$routeChangeStart', function (event, currRoute, prevRoute) {
            var logged = AuthService.isAuthenticated();
            //check if the user is going to the login page
            // i use ui.route so not exactly sure about this one but you get the picture

            if (!logged) {
                // event.preventDefault();
                $location.path('/login');

            }
        });
    }])
    .factory('localStorage', ['$window', function ($window) {
        if ($window.localStorage) {
            return $window.localStorage;
        }
        throw new Error('Local storage support is needed');
    }])
    .service('Session', function ($log, localStorage) {
        // Instantiate data when service
        // is loaded
        this._user = JSON.parse(localStorage.getItem('session.user'));
        this._accessToken = localStorage.getItem('session.accessToken');

        this.getUser = function () {
            return this._user;
        };

        this.setUser = function (user) {
            this._user = user;
            localStorage.setItem('session.user', JSON.stringify(user));
            return this;
        };

        this.getAccessToken = function () {
            var t = this._accessToken;
            if (t === "null") {
                t = null;
            }
            return t;
        };

        this.setAccessToken = function (token) {
            this._accessToken = token;
            localStorage.setItem('session.accessToken', token);
            return this;
        };

        /**
         * Destroy session
         */
        this.destroy = function destroy() {
            this.setUser(null);
            this.setAccessToken(null);
        };
    })
    .factory('AuthService', ['$http', '$log', 'Session', function ($http, $log, Session) {
        var authService = {};

        authService.login = function (credentials) {
            console.log("Got creds:", credentials);
            return $http
                .post('/login', credentials)
                .then(function (res) {
                    if (res.data.user) {
                        Session.setUser(res.data.user);
                        if (res.data.user.token) {
                            Session.setAccessToken(res.data.user.token);
                        }
                    }

                    return res.data;
                }, function (err) {
                    console.log(err)
                });
        };

        authService.isAuthenticated = function () {
            var t = Session.getAccessToken();
            if (t === "null") {
                t = null;
            }
            return t !== null;
        };
        authService.logout = function () {
            return Session.destroy();
        }

        return authService;
    }])