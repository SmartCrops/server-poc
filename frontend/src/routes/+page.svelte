<script lang="ts">
    import Chart from "$lib/Chart.svelte";
    import { Table, tableMapperValues } from "@skeletonlabs/skeleton";
    import type { TableSource } from "@skeletonlabs/skeleton";

    const sourceData = [
        { position: 1, name: "Hydrogen", weight: 1.0079, symbol: "H" },
        { position: 2, name: "Helium", weight: 4.0026, symbol: "He" },
        { position: 3, name: "Lithium", weight: 6.941, symbol: "Li" },
        { position: 4, name: "Beryllium", weight: 9.0122, symbol: "Be" },
        { position: 5, name: "Boron", weight: 10.811, symbol: "B" },
    ];
    const tableSimple: TableSource = {
        // A list of heading labels.
        head: ["Position", "Name", "Symbol", "Weight"],
        // The data visibly shown in your table body UI.
        body: tableMapperValues(sourceData, [
            "position",
            "name",
            "symbol",
            "weight",
        ]),
    };
</script>

<div class="container mx-auto p-8 space-y-8">
    <h1>Welcome to <strong>SmartCrops</strong></h1>
    <div class="grid grid-cols-2 gap-2">
        <div class="card text-center">
            <header class="card-header pb-4">
                <h2>Chart Example</h2>
            </header>
            <Chart />
        </div>
        <div class="card">
            <header class="card-header text-center pb-4">
                <h2>Table Example</h2>
            </header>
            <Table source={tableSimple} interactive={true} />
        </div>
    </div>
</div>
