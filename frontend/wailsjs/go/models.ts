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
	export class user_model {
	    id: number;
	    name: string;
	    ip: string;
	
	    static createFrom(source: any = {}) {
	        return new user_model(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.ip = source["ip"];
	    }
	}

}

