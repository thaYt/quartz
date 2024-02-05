<script lang="ts">
  import { onMount } from "svelte";
  import Navbar from "./Navbar.svelte";
  import Playerlist from "./Playerlist.svelte";
  import Settings from "./Settings.svelte";

  onMount(() => {
    console.log("App mounted");
    $state = 0;
  });
  
  import { writable } from "svelte/store";

  enum State {
    Playerlist,
    Settings,
  }

  const initialState: State = State.Playerlist;

  export const stateStore = writable(0);
  export const update = (value: State) => stateStore.set(value);
  export const reset = () => stateStore.set(initialState);

  export const state = {
    subscribe: stateStore.subscribe,
    update,
    reset,
  };
</script>

<main>
  <div class="nav">
    <Navbar />
  </div>

  <hr />

  {#if $state === State.Playerlist}
    <Playerlist />
  {:else if $state === State.Settings}
    <Settings />
  {/if}
</main>

<style lang="scss">
  .nav {
    height: 40px;
  }

  hr {
    margin: 0;
    padding: 0;
    border-color: rgba(15, 15, 15, 0.3);
  }
</style>
