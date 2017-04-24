var app = angular.module('gopreso.login.controller', [
    'ngMaterial',
    'ngAnimate',
    'ngAria',
    'ngMessages'
])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/login', {
            templateUrl: '/static/views/login.html',
            controller: 'LoginController',
            name: "login"
        });
    }])
    .controller('LoginController', function ($scope, $http, authService) {
        $scope.vm = {
            formData: {
                email: 'admin',
                password: 'admin'
            },
            submit: function () {
                console.log($scope.vm.formData)
                $http({
                    method: "POST",
                    url: "/api/login",
                    data: {
                        username: $scope.vm.formData.email,
                        password: $scope.vm.formData.password
                    }
                }).then(function (data) {
                    console.log("Successful Login")

                    authService.loginConfirmed()
                }, function (err) {
                    console.log(err)
                })
            }
        };

    });