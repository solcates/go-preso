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
    .controller('LoginController', function ($scope, $http) {
        $scope.vm = {
            formData: {
                email: 'hello@patternry.com',
                password: 'foobar'
            },
            submit: function () {
                console.log($scope.vm.formData)
                $http({
                    method: "POST",
                    url: "/login",
                    data: {
                        username: $scope.vm.formData.email,
                        password: $scope.vm.formData.password
                    }
                })
            }
        };

    });