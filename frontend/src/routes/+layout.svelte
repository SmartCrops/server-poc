<script>
  import "@skeletonlabs/skeleton/themes/theme-skeleton.css";
  import "@skeletonlabs/skeleton/styles/all.css";
  import "../app.postcss";

  import {
    cubes,
    home,
    barChart,
    questionCircle,
    userCircle,
  } from "svelte-awesome/icons";

  import { AppShell, AppBar } from "@skeletonlabs/skeleton";
  import IconLink from "$lib/IconLink.svelte";
</script>

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
  <!-- Header -->
  <svelte:fragment slot="header">
    <AppBar>
      <svelte:fragment slot="lead">
        <h1>SmartCrops</h1>
      </svelte:fragment>
      <svelte:fragment slot="trail">
        <a class="btn btn-sm" href="/" target="_blank" rel="noreferrer">
          Log Out
        </a>
      </svelte:fragment>
    </AppBar>
  </svelte:fragment>

  <!-- Sidebar -->
  <svelte:fragment slot="sidebarLeft">
    <!-- Insert the list: -->
    <nav class="list-nav">
      <ul>
        <li><IconLink href="/" icon={home} text="Home" /></li>
        <li><IconLink href="/" icon={barChart} text="Dashboard" /></li>
        <li>
          <IconLink
            href="/datacollectors"
            icon={cubes}
            text="Data Collectors"
          />
        </li>
        <li><IconLink href="/" icon={cubes} text="Pump Controllers" /></li>
        <li><IconLink href="/" icon={userCircle} text="My Account" /></li>
        <li><IconLink href="/" icon={questionCircle} text="Help" /></li>
      </ul>
    </nav>
    <!-- --- -->
  </svelte:fragment>

  <!-- Page Content Slot -->
  <slot />
</AppShell>
