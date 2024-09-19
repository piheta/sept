export namespace models {
	
	export class Message {
	    id: number;
	    chat_id: number;
	    user_id: number;
	    content: string;
	    created_at: string;
	    signature: string;
	
	    static createFrom(source: any = {}) {
	        return new Message(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.chat_id = source["chat_id"];
	        this.user_id = source["user_id"];
	        this.content = source["content"];
	        this.created_at = source["created_at"];
	        this.signature = source["signature"];
	    }
	}
	export class User {
	    id: string;
	    username: string;
	    ip: string;
	    avatar: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.ip = source["ip"];
	        this.avatar = source["avatar"];
	    }
	}

}

