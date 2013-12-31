


var Test = function() {

	console.log('test model constructor function');
	var self = this;

	this.title = 'Loading...';

	jQuery.ajax({
		type: 'GET',
		url: 'http://echo.jsontest.com/title/My%20App%20Matt',
		success: function(data) {
			self.title = decodeURIComponent(data.title);
		},
		dataType: 'jsonp',
		crossDomain: true
	});
};




module.exports = Test;
