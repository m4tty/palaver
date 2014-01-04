var page = require('./libs/page.js');
//var ractive = require('./views/test.js');



console.log('app.js');
//page.base('/static/index.html');

page('/static/index.debug.html?:section', section);

page('/static/:section',section);

page('/static/index.debug.html#profile', function() {
//	ractive.set({title: 'win'});

	console.log('profile');
});

page('*', function(ctx, next) {
//	ractive.set({title: 'win'});
	console.log('path * :',ctx.path);
	console.log('hash:',ctx.hash);
	console.log('ctx',ctx);
	console.log('app js page nav * function');
});

function section(ctx,next) {
	console.log('section');
	console.log('path:',ctx.path);
	console.log('hash:',ctx.hash);
}


page();