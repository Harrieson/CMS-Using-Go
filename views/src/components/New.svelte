<script>
  import { Modals, closeModal, openModal } from "svelte-modals";
  import Modal from "./Modal.svelte";

  async function updateGoly(data) {
    const json = {
      redirect: data.redirect,
      goly: data.goly,
      random: data.random,
    };

    await fetch("http://localhost:3000/redirect", {
      method: "POST",
      header: { "Content-Type": "application/json" },
      body: JSON.stringify(json),
    }).then((res) => {
      console.log(res);
    });
  }

  function handleOpen() {
    openModal(Modal, {
      title: "Create New Goly Link",
      send: updateGoly,
      redirect: "",
      goly: "",
      random: false,
    });
  }
</script>

<button on:click={handleOpen}>New Goly</button>

<style>
  button {
    background-color: green;
    color: white;
    font-weight: bold;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
    cursor: pointer;
  }
</style>
