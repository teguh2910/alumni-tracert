<script>
	import { Router, Route } from 'svelte-routing';
	import ProtectedRoute from './ProtectedRoute.svelte';
	import Toast from './components/Toast.svelte'
	import Home from './routes/Home.svelte';
	import Dashboard from './routes/Dashboard.svelte';
	import Kuisioner from './routes/Kuisioner.svelte';
	import KuisionerForm from './routes/KuisionerForm.svelte'
	import AlumniRegistration from './routes/AlumniRegistration.svelte'
	import AppraiserRegistration from './routes/AppraiserRegistration.svelte'
	import DetailAlumni from './routes/AlumniDetail.svelte';
	import Upload from './routes/Upload.svelte';
	import QrIjazahLandingPage from './routes/ijazah/[id].svelte';
	import QrTranscriptLandingPage from './routes/transcript/[id].svelte';
	import Login from './routes/Login.svelte';
	import ListAlumni from './routes/list-alumni.svelte';
	import ELegalisir from './routes/e-legalisir.svelte';
	import { PATH_URL } from './helper/path';
	import Cdc from './routes/Cdc.svelte';
	import SapaAlumni from './routes/SapaAlumni.svelte';
	import TentangKami from './routes/TentangKami.svelte';
	import LegalisirDetail from './routes/LegalisirDetail.svelte'
	import Ijazah from './routes/Ijazah.svelte'
  import ForgotPassword from './routes/ForgotPassword.svelte'
  import ResetPassword from './routes/ResetPassword.svelte'
</script>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
	@import url('https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css');
</style>

<Router>
	<Route path={PATH_URL.BASE} component={Home} />
	<Route path={PATH_URL.LOGIN} component={Login} />
	<Route path={PATH_URL.KUISIONER} component={Kuisioner} />
	<Route path={PATH_URL.CDC} component={Cdc} />
	<Route path={PATH_URL.SAPA_ALUMNI} component={SapaAlumni} />
	<Route path={PATH_URL.TENTANG_KAMI} component={TentangKami} />
	<Route path={PATH_URL.ALUMNI_REGISTRATION} component={AlumniRegistration} />
	<Route path={PATH_URL.APPRAISER_REGISTRATION} component={AppraiserRegistration} />
  <Route path={PATH_URL.FORGOT_PASS} component={ForgotPassword} />
  <Route path={PATH_URL.RESET_PASS} let:params>
    <ResetPassword id="{params.id}"/>
  </Route>
	<Route path={PATH_URL.QR_IJAZAH_LANDING_PAGE} let:params>
		<QrIjazahLandingPage id="{params.id}"/>
	</Route>
	<Route path={PATH_URL.QR_TRANSCRIPT_LANDING_PAGE} let:params>
		<QrTranscriptLandingPage id="{params.id}"/>
	</Route>
	<ProtectedRoute path={PATH_URL.DASHBOARD} component={Dashboard} />
	<ProtectedRoute path={PATH_URL.KUISIONER_FORM} component={KuisionerForm} />
	<ProtectedRoute path={PATH_URL.IJAZAH} component={Ijazah} />
	<ProtectedRoute path={PATH_URL.UPLOAD_IJAZAH} component={Upload} />
	<ProtectedRoute path={PATH_URL.ADMIN_ALUMNI} component={ListAlumni} />
	<ProtectedRoute path={PATH_URL.ADMIN_E_LEGALISIR} component={ELegalisir} />
	<ProtectedRoute path={PATH_URL.ADMIN_ALUMNI_DETAIL} component={DetailAlumni} />
	<ProtectedRoute path={PATH_URL.ADMIN_LEGALISIR_DETAIL} component={LegalisirDetail} />
  <ProtectedRoute path={PATH_URL.CHANGE_PASS} component={ResetPassword} />
</Router>
<Toast />