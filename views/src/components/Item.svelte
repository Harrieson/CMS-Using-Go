<script>
  import Card from "./Card.svelte";
  import { Modals, closeModal, openModal } from "svelte-modals";
  export let goly;

  async function update(data) {
    const json = {
      redirect: data.redirect,
      goly: data.goly,
      random: data.random,
      id: goly.id,
    };
  }

  function handleOpen(goly) {
    openModal(Modals);
  }
</script>

<Card>
  <p>Goly: http://localhost:3000/r/{goly.goly}</p>
  <p>Redirect: {goly.redirect}</p>
  <p>Clicked: {goly.clicked}</p>
  <button class="update" on:click={handleOpen}>Update</button>
  <button class="delete">Delete</button>
</Card>
<Modals>
  <div on:click={closeModal} />
</Modals>

<style>
  button {
    color: white;
    font-weight: bolder;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
    cursor: pointer;
  }

  .update {
    background-color: yellowgreen;
  }

  .delete {
    background-color: red;
  }
</style>
