<script>
  import { Route } from 'svelte-routing'
  import accessDenied from './routes/accessDenied.svelte'
  // import Cookies from 'js-cookie'
  import { token } from './stores/token.js';

  export let path
  export let component
  $: isAuthenticated = $token;
</script>

{#if isAuthenticated}
  <Route path={path} component={component} />
{:else}
  <Route path={path} component={accessDenied} />
{/if}