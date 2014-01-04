


var TopNavModel = function() {

	console.log('test model constructor function');
	var self = this;

	this.screenName = 'log in';
	this.avatarUrl = '';
	this.useGravatar = false;
	this.isSignedIn = false;
	this.loginLocation = "/login/google";
	this.currentLocation = window.location.pathname;

	jQuery.ajax({
		type: 'GET',
		url: 'http://localhost:8080/users/current',
		success: function(data) {
			console.log('success callback');
			self.screenName = decodeURIComponent(data.ScreenName);
			//console.log(data);
			self.avatarUrl = decodeURIComponent(data.AvatarUrl);
			self.isSignedIn = true;
			self.useGravatar = false;
		},
		dataType: 'json',
		crossDomain: true
	});
};





module.exports = TopNavModel;
