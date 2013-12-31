


var TopNavModel = function() {

	console.log('test model constructor function');
	var self = this;

	this.screenName = 'log in';
	this.avatarUrl = '';
	this.useGravatar = false;
	this.isSignedIn = false;
	jQuery.ajax({
		type: 'GET',
		url: 'http://localhost:8080/users/current',
		success: function(data) {
			console.log('success callback');
			self.screenName = decodeURIComponent(data.ScreenName);
			//console.log(data);
			self.avatarUrl = decodeURIComponent(data.AvatarUrl);
			self.isSignedIn = true
		},
		dataType: 'json',
		crossDomain: true
	});
};





module.exports = TopNavModel;
