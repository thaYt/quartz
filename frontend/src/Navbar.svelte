<script lang="ts">
  import { onMount } from "svelte";
  import { fly } from "svelte/transition";
  import { WindowMinimise, Quit, EventsEmit } from "../wailsjs/runtime";
  import { GetVersion } from "../wailsjs/go/app/App";

  import Close from "./assets/images/nav/Close.svelte";
  import Minimize from "./assets/images/nav/Minimize.svelte";

  import quartz from "./assets/images/nav/Nether_Quartz_JE2_BE2.webp";
  import Search from "./assets/images/nav/Search.svelte";

  onMount(async () => {
    document.querySelector(".navbar-version").textContent =
      `v${await GetVersion()}`;
  });

  const toggleSearch = () => (searching = !searching);

  function handleSearch() {
    EventsEmit("SendPlayer", searchInput);
    searchInput = "";
  }

  let searchInput = "";
  let searching = false;
</script>

<div class="navbar draggable">
  <div class="navbar-left">
    <img src={quartz} width="30" height="30" alt="quartz" />

    <div class="navbar-title">quartz</div>

    <div class="navbar-version" />
  </div>
  <!-- add close, minimize buttons -->
  <div class="navbar-buttons">
    <div class="search">
      {#if searching}
        <button class="btn" id="search-btn" on:click={toggleSearch}>
          <Search stroke="white" />
        </button>
        <div
          class="search-bar undraggable"
          transition:fly={{ x: 50, duration: 200 }}
        >
          <input
            type="text"
            placeholder="search player..."
            bind:value={searchInput}
            on:keydown={(e) => e.key === "Enter" && handleSearch()}
          />
        </div>
      {:else}
        <button class="btn" id="search-btn" on:click={toggleSearch}>
          <Search stroke="white" />
        </button>
      {/if}
    </div>
    <button class="btn" id="min-btn" on:click={WindowMinimise}>
      <Minimize stroke="white" />
    </button>
    <button class="btn" id="close-btn" on:click={Quit}>
      <Close stroke="white" />
    </button>
  </div>
</div>

<style lang="scss">
  .navbar {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 40px;
    line-height: 40px;
    background-color: rgba(15, 15, 15, 0.3);
    z-index: 100;
    user-select: none;
    align-items: start;
  }

  .navbar-left {
    position: absolute;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    font-size: 24px;
  }

  .navbar-buttons {
    position: absolute;
    right: 0;
    height: 40px;
    line-height: 40px;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
  }

  .draggable {
    --wails-draggable: drag;
  }

  .undraggable {
    --wails-draggable: none;
  }

  .navbar .btn {
    border: none;
    background-color: rgba(15, 15, 15, 0);
    cursor: pointer;
    transition: background-color 0.2s ease-in-out;
    padding-top: 3px;
    padding-bottom: 0px;
    padding-left: 10px;
    padding-right: 10px;
    margin: 0px;
  }

  .navbar .btn:hover {
    background-color: rgba(15, 15, 15, 0.2);
  }

  .navbar img {
    margin-left: 10px;
    margin-right: 10px;
    padding-top: 3px;
  }

  .navbar-version {
    margin-left: 10px;
    color: gray;
  }

  .search {
    display: flex;
    align-items: center;
  }

  .search-bar {
    border-radius: 4px;
    width: 200px;
    transition: width 0.2s ease-in-out;
  }

  .search-bar input {
    background-color: rgba(15, 15, 15, 0.3);
    width: 100%;
    height: 25px;
    border: none;
    border-radius: 4px;
    font-size: 15;
    color: white;
  }

  .search-bar input:focus {
    outline: none;
  }
</style>
