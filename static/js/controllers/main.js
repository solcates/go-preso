var app = angular.module('gopreso.main.controller', [
    'ngRoute'
])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/', {
            templateUrl: '/static/views/main.html',
            controller: 'MainController',
            controllerAs: "main",
            name: "main"
        });
    }])
    .controller('MainController', function ($scope, $http, $log) {
        $scope.presos = [];
        $scope.selected_preso = [];
        $http({
            url: "/api/presos",
            method: "GET"
        }).then(function(data){
            $scope.presos = data.data
            $log.log($scope.presos)
        })
    });