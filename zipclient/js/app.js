'use strict';

var myApp = angular.module('GoClient', ['ngSanitize', 'ui.router']);

//configure routes
myApp.config(['$stateProvider', '$urlRouterProvider', function ($stateProvider, $urlRouterProvider) {
	$stateProvider
		.state('home', {
			url: '/home',
			templateUrl: 'partials/home.html',
			controller: 'ZipsCtrl'
		})
	$urlRouterProvider.otherwise('/home');
}]);

//For movie list
myApp.controller('ZipsCtrl', ['$scope', '$http', function ($scope, $http) {

  $scope.search = function() {
  	$http.get('http://localhost:4200/zips/' + $scope.city).then(function (response) {
		var data = response.data;
		$scope.zips = data;
  });
  }
}]);