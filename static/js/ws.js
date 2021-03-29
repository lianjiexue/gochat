// ws.js
	var ws = new WebSocket("ws://127.0.0.1:8080/ws");
	
	ws.onclose = function(){
		console.log("链接关闭")
	}

	ws.onopen = function(evt){
		console.log(evt.data)
		//链接成功以后进行绑定
	}

	ws.onmessage = function(evt){
		console.log(evt.data)
	}