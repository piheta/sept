export namespace main {
	
	export class server_model {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new server_model(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}

}

export namespace models {
	
	export class User_model {
	    id: number;
	    user_id: string;
	    username: string;
	    ip: string;
	    avatar: string;
	
	    static createFrom(source: any = {}) {
	        return new User_model(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.user_id = source["user_id"];
	        this.username = source["username"];
	        this.ip = source["ip"];
	        this.avatar = source["avatar"];
	    }
	}

}

