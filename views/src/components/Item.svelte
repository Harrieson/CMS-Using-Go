<script>
  import Card from "./Card.svelte";
  import Modal from "./Modal.svelte";
  import { Modals, closeModal, openModal } from "svelte-modals";
  export let goly;
  let ShowCard = true;

  async function update(data) {
    const json = {
      redirect: data.redirect,
      goly: data.goly,
      random: data.random,
      id: goly.id,
    };

    await fetch("http://localhost:3000/redirect", {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(json),
    }).then((res) => {
      console.log(res);
    });
  }

  function handleOpen(goly) {
    openModal(Modal, {
      title: "Update Goly Link",
      send: update,
      goly: goly.goly,
      redirect: goly.redirect,
      random: goly.random,
    });
  }

  async function deleteGoly() {
    if (
      confirm(
        "Are you sure you want to delete this Goly link (" + goly.goly + ")?"
      )
    ) {
      await fetch("http://localhost:3000/goly/" + goly.id, {
        method: "DELETE",
      }).then((res) => {
        ShowCard = false;
        console.log(res);
      });
    }
  }
</script>

{#if ShowCard}
  <Card>
    <p>Goly: http://localhost:3000/r/{goly.goly}</p>
    <p>Redirect: {goly.redirect}</p>
    <p>Clicked: {goly.clicked}</p>
    <button class="update" on:click={handleOpen}>Update</button>
    <button class="delete" on:click={deleteGoly}>Delete</button>
  </Card>
{/if}
<Modals>
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div on:click={closeModal} slot="backdrop" class="backdrop" />
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
  .backdrop {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    background: rgb(255, 255, 255);
  }
</style>
