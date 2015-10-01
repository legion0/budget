angular.module('entryDirective', [])
.directive('budgetEntry', function() {
	return {
		restrict: 'E',
		scope: false,
		templateUrl: '/components/entry/view.html'
	}
})
.controller('EntryController', ['$scope', 'Entry', function($scope, Entry) {  
    $scope.entry = new Entry();
}]);;

// model
app.factory('Entry', ['$http', function($http) {  
	function Entry(entryData) {
		this.id = 0;
		this.userId = 0;
		this.name = undefined;
		this.price = undefined;
		this.when = new Date();
		this.when.setSeconds(0);
		this.when.setMiliseconds(0);
		if (entryData) {
			this.setData(entryData);
		}
		// Some other initializations related to entry
	};
	Entry.prototype = {
		setData: function(entryData) {
			angular.extend(this, entryData);
		},
		save: function() {
			console.log(this);
			$http.put('/entry/' + this.id, this);
		}
	};
	return Entry;
}]);

