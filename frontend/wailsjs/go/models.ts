export namespace models {
	
	export class Chat {
	    id: string;
	    name: string;
	    avatar: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Chat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.avatar = source["avatar"];
	        this.created_at = source["created_at"];
	    }
	}
	export class Message {
	    id: number;
	    chat_id: string;
	    user_id: string;
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
	    public_key: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.ip = source["ip"];
	        this.avatar = source["avatar"];
	        this.public_key = source["public_key"];
	    }
	}

}

