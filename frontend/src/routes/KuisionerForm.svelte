<script>
  import { onMount, onDestroy } from 'svelte'
  import { navigate } from 'svelte-routing'
  import { globalHistory } from 'svelte-routing/src/history';
  import Cookies from 'js-cookie'

  import { Images } from '../helper/images'
  import { PATH_URL } from '../helper/path'
  import { notifications } from '../helper/toast'
  
  import { TracertServicePromiseClient } from '../../proto/tracert_service_grpc_web_pb'
  import { QuestionGroupListInput, QuestionGroupList } from '../../proto/question_group_message_pb'
  import { UserAnswer, Tracer } from '../../proto/user_answer_message_pb'
  import QuestionService from '../services/question'
  import UserAnswerService from '../services/user_answer'
  import errorServiceHandling from '../helper/error_service'
  
  let groups = [1];
  onMount(() => {
    if (Cookies.get('usertype') == 2) {
      groups = [6];
    }
  })
  const userAnswer = [];
  let tracerId = null;
  
  const questionGroupListInputProto = new QuestionGroupListInput()
  let questionList = new QuestionGroupList()
  
  async function questionListCall(){
    var deps = {
			proto: {
				TracertClient: TracertServicePromiseClient
			}
		}
    questionGroupListInputProto.setQuestionGroupIdList(groups)
    const question = new QuestionService(deps, questionGroupListInputProto)
    return await question.list()
	}

  async function tracerCreateCall(){
    var deps = {
      proto: {
        TracertClient: TracertServicePromiseClient
      }
    }

    const userId = Cookies.get('userid')

    const tracerProto = new Tracer();

    tracerProto.setUserId(userId);
    const tracerService = new UserAnswerService(deps, tracerProto)

    return await tracerService.tracer()
  }

  async function userAnswerCall(){
    var deps = {
			proto: {
				TracertClient: TracertServicePromiseClient
			}
		}

    if (!tracerId) {
      let tracerResponse = await tracerCreateCall();
      tracerId = tracerResponse.getId();
    }
    
    let promises = [];
    userAnswer.forEach((answer, questionId) => {
      console.log(`answer`, answer)
      const userAnswerProto = new UserAnswer()
      // userAnswerProto.setUserId(userId)
      userAnswerProto.setTracerId(tracerId)
      userAnswerProto.setQuestionId(questionId)
      if (Array.isArray(answer)) {
        userAnswerProto.setAnswer(JSON.stringify(answer));
      } else {
        userAnswerProto.setAnswer(JSON.stringify(answer.text));
      }
      
      const userAnswerService = new UserAnswerService(deps, userAnswerProto)
      promises.push(userAnswerService.answer())
    })

    return Promise.all(promises)
	}

  onMount(async () => {
		try {
      questionList = await questionListCall()
      // let questionGroupList = questionList.getQuestionGroupList();
      // console.log(`questionList`, questionGroupList)
      // questionGroupList.forEach(group => {
      //   console.log(`group`, group.getTitle())
      // })
    } catch(e) {
      errorServiceHandling(e)
      if (Cookies.get('token') == null) {
        location = PATH_URL.LOGIN  
      } 
      notifications.danger(e.message)
    }
	})

  const changeAnswer = (event, questionId, answerTitle, isMultiple) => {
    const answer = {
      id: event.target.value,
      text: answerTitle,
    };

    
    if (isMultiple) {
      if (!userAnswer[questionId]) {
        let temp = [];
        temp.push(answerTitle)
        userAnswer[questionId] = temp;
      } else {
        let temp = userAnswer[questionId];
        let value = answerTitle;
        // let isChecked = event.target.checked;
        if (temp.indexOf(answerTitle) === -1) {
          temp.push(answerTitle);
        } else {
          temp.splice(temp.indexOf(answerTitle), 1);
        }
        userAnswer[questionId] = temp;
      }
    } else {
      // console.log(questionId, event.target.value)
      userAnswer[questionId] = answer
    }
  }

  const validateAnswer = () => {
    let result = true;
    questionList.getQuestionGroupList().forEach(group => {
      group.getQuestionList().forEach(question => {
        console.log(question.getId(), userAnswer[question.getId()])
        if(!userAnswer[question.getId()]) {
          result = false
        } 
      })
    })
    return result
  }

  const lanjutkan = async () => {
    try {
      if (!validateAnswer()) {
        throw { message: "silahkan jawab kuisioner terlebih dahulu"}
      } 

      const answer = await userAnswerCall()
      if (groups.length === 1 && groups[0] === 1) {
        if (userAnswer[1].id !== 5) {
          groups = [(parseInt(userAnswer[1].id)+1)]
        } 
        questionList = await questionListCall()
      } else {
        if (Cookies.get('usertype') == 2) {
          navigate(PATH_URL.DASHBOARD, { replace: true })
        } else {
          navigate(PATH_URL.UPLOAD_IJAZAH, { replace: true })
        }
      }
    } catch(e) {
      console.log(`e`, e)
      errorServiceHandling(e)
      if (Cookies.get('token') == null) {
        location = PATH_URL.LOGIN  
      } 
      notifications.danger(e.message)
    }
  }

  let rangeQuestion = true;

  const isCheckedRadio = (question, questionOption) => {
    let flag = false

    if (userAnswer[question.getId()] && userAnswer[question.getId()].id) {
      flag = userAnswer[question.getId()].id == questionOption.getId()
    }

    return flag;
  }

</script>

<div class="flex flex-wrap w-full h-full">
  <div class="flex w-full p-4 align-center">

    <main class="max-w-3xl px-4 mx-auto my-24 sm:mt-12 sm:px-6 md:mt-16 lg:my-24 lg:px-8">
      <div class="sm:text-center lg:text-left">
        
        <div class="sticky top-0 pt-6 bg-white">
          <a href="/" class="flex items-center mb-8">
            <i class="mr-4 fas fa-arrow-left"></i>
            <p class="text-base">Kembali ke halaman utama</p>
          </a>
          <img class="object-cover w-64 h-full mb-4" src={Images.logo_poltekkes} alt="">
          <h1 class="text-3xl font-extrabold tracking-tight text-gray-900 sm:text-3xl md:text-3xl">
            <span class="block xl:inline">KUISIONER TRACER STUDY/PENGGUNA ALUMNI</span>
          </h1>
          
          <hr class="my-8 md:min-w-full" />
        </div>

        <div class="">
          {#each questionList.getQuestionGroupList() as group}
            
          <h2 class="block mb-8 font-bold text-blue-700">{group.getTitle()}</h2>

            {#each group.getQuestionList() as question}
              <div class="mb-12">
                
                <p class="mb-4 text-xl font-semibold text-black">{question.getTitle()}</p>

                <div class={question.getMinimumValue() && question.getMaximumValue() ? "flex justify-between mt-6 items-center" : "mt-2 space-y-4"}>
                  
                  {#if question.getQuestionType() === 1}
                    <input on:change="{(event) => changeAnswer(event, question.getId(), event.target.value)}" type="text" name="first-name" id="first-name" autocomplete="given-name" class="block w-full px-4 py-2 border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-m">
                  {:else if  question.getQuestionType() === 2}

                    {#if question.getMinimumValue() && question.getMaximumValue() }
                      <span>{question.getMinimumValue()}</span>
                    {/if}

                    {#each question.getQuestionOptionList() as questionOption, i}
                      <div class="flex items-center">
                        <input checked={ isCheckedRadio(question, questionOption)} on:change="{(event) => changeAnswer(event, question.getId(), questionOption.getTitle())}" id={`radio-${question.getId()}-${i}`} name={`radio-${question.getId()}`} type="radio" value="{questionOption.getId()}" class="w-4 h-4 text-indigo-600 border-gray-300 focus:ring-indigo-500">
                        <!-- <input checked={ userAnswer[question.getId()] == questionOption.getId()} on:change="{(event) => changeAnswer(event, question.getId(), questionOption.getTitle())}" id={`radio-${question.getId()}-${i}`} name={`radio-${question.getId()}-${i}`} type="radio" value="{questionOption.getId()}" class="w-4 h-4 text-indigo-600 border-gray-300 focus:ring-indigo-500">                         -->
                        <div class="flex items-center w-full">
                          <label for={`radio-${question.getId()}-${i}`} class="block max-w-xl ml-3 text-sm font-medium text-gray-700 cursor-pointer w-max min-w-max">
                            {questionOption.getTitle()} 
                          </label>
                          {#if questionOption.getIsNeedEssay()}
                            <input type="text" class="block w-full px-4 py-2 ml-2 border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-m">
                          {/if}
                        </div>

                      </div>
                    {/each}

                    {#if question.getMinimumValue() && question.getMaximumValue() }
                        <span>{question.getMaximumValue()}</span>
                    {/if}
                    
                  {:else if  question.getQuestionType() === 3}
                    {#each question.getQuestionOptionList() as questionOption, i}
                      <div class="flex items-start">
                        <div class="flex items-center h-5">
                          <input value={questionOption.getId()} on:change="{(event) => changeAnswer(event, question.getId(), questionOption.getTitle(), true)}" id={`check-${question.getId()}-${i}`} name={`check-${question.getId()}-${i}`} type="checkbox" class="w-4 h-4 text-indigo-600 border-gray-300 rounded focus:ring-indigo-500">
                        </div>
                        <div class="ml-3 text-sm">
                          <label for={`check-${question.getId()}-${i}`} class="font-medium text-gray-700 cursor-pointer">{questionOption.getTitle()}</label>
                        </div>
                      </div>
                    {/each}
                  {/if}
                </div>
              </div>
            {/each}
          {:else}
            <!-- this block renders when photos.length === 0 -->
            <p>loading...</p>
          {/each}
        
          
          <div class="px-4 py-3 mt-10 text-right bg-gray-50 sm:px-6">
            <button on:click="{lanjutkan}" class="inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-indigo-600 border border-transparent rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Lanjutkan
            </button>
          </div>
        </div>

      </div>
    </main>

  </div>
</div>