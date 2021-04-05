// ws.js
	var ws = new WebSocket("ws://127.0.0.1:8080/ws");
	
	ws.onclose = function(){
		console.log("链接关闭")
	}

	ws.onopen = function(evt){
		console.log(evt.data)
		//链接成功以后进行绑定
		bind()
	}

	ws.onmessage = function(evt){
		if(evt.data.type == "message") {
			console.log("收到消息"+evt.data.content)
		} else {
			console.log(evt.data)
		}
		
	}
		
	function bind(){
		var obj = {};
			obj.type = "bind";
			obj.FromId = parseInt(localStorage.getItem("uid"));
		ws.send(JSON.stringify(obj));
	}