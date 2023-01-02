<script>
  import Card from "./Card.svelte";
</script>

<Card>
  <p>Goly: http://localhost:3000/r/ {"asdfg"}</p>
  <p>Redirect: {"redirect"}</p>
  <p>Clicked:</p>
  <button class="update">Update</button>
  <button class="delete">Delete</button>
</Card>

<style>
  button {
    color: white;
    font-weight: bolder;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
  }

  .update {
    background-color: yellowgreen;
  }

  .delete {
    background-color: red;
  }
</style>
