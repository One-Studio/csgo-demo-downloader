export namespace config {
	
	export class CFG {
	    version: string;
	    demoDir: string;
	    useExternel: boolean;
	    autoDownload: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CFG(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.demoDir = source["demoDir"];
	        this.useExternel = source["useExternel"];
	        this.autoDownload = source["autoDownload"];
	    }
	}

}

