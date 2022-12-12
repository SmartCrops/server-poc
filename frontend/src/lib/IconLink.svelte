<script lang="ts">
    import Icon from "svelte-awesome/components/Icon.svelte";

    export let href: string;
    export let text: string;
    export let icon: any;
</script>

<a {href}>
    <span>
        <Icon data={icon} />
    </span>
    <span> {text} </span>
</a>
