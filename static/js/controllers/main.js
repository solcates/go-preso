var app = angular.module('gopreso.main.controller', [
    'ngRoute'
])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/login', {
            templateUrl: '/static/views/main.html',
            controller: 'MainController',
            controllerAs: "main",
            name: "main"
        });
    }])
    .controller('MainController', function ($scope, $http) {

    });