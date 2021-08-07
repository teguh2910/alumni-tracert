import { HOST_URL, APP_ENV } from '../env'
export default class{
  constructor(deps){
      this.proto = deps.proto
      this.client = new deps.proto.TracertClient(HOST_URL, null, null)
      
      if (APP_ENV === 'development') {
        const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {})
        enableDevTools([
          this.client,
        ]);
      }
  }
}