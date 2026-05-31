import Tracert from './tracert';
import Cookies from 'js-cookie'
import { EmptyMessage } from '../../proto/generic_message_pb'

export default class extends Tracert{
  constructor(proto, UserAnswer) {
    super(proto)
    this.req = UserAnswer
  }
  answer (){
      return this.client.userAnswerCreate(this.req, {'token':Cookies.get('token')}).then(userAnswer=>{
          return userAnswer
      })
  }
  list() {
    return this.client.getTrace(this.req, {'token':Cookies.get('token')});
  }
  tracer (){
      return this.client.tracerCreate(this.req, {'token':Cookies.get('token')}).then(createdTracer=>{
          return createdTracer
      })
  }
  getMyAnswers() {
    return this.client.getMyAnswers(new EmptyMessage(), {'token': Cookies.get('token')});
  }
}