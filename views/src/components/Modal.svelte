<script>
  import { closeModal } from "svelte-modals";
  export let title;
  export let isOpen;
  export let redirect;
  export let goly;
  export let random;
  export let send;

  let data = {
    redirect: redirect,
    goly: goly,
    random: random,
  };
</script>

<div role="dialog" class="modal">
  <div class="contents">
    <h2>{title}</h2>
    <label for="redirect">Redirect To:</label>
    <input type="text" name="redirect" id="redirect" bind:value={redirect} />

    <label for="goly">Goly</label>
    <input type="text" name="goly" id="goly" bind:value={data.goly} />
  </div>
</div>
