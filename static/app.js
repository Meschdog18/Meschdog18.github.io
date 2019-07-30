new Vue({
	el:'#vue-app',
	data: {
		health: 'hey',
		testy: null
	},
	mounted: function() {
    

    var urlParams = new URLSearchParams(window.location.search);
  		    	console.log(urlParams.get('test'));
  		    	var x = urlParams.get('test');
  		    	console.log(x)
  		    	 this.testy = x; 
  }
});