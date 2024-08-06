<script>
    // @ts-nocheck
    import Sidebar from "./components/Sidebar/Sidebar.svelte";
    import Footer from "./components/Footer.svelte";
    import { onDrag } from "./components/Sidebar/SidebarResizer";
    import Header from "./components/Header.svelte";
    let width = 205;
    let resizeWidth = width;
    const minWidth = 205;
    const maxWidth = () => window.innerWidth / 2;

    $: console.log(resizeWidth);

    function handleDrag(delta) {
        const newWidth = width + delta;
        resizeWidth = Math.max(minWidth, Math.min(maxWidth(), newWidth));
    }
</script>

<main class="h-full">
    <div class="wails-drag w-full h-6 m-0"></div>
    <div class="flex w-full h-[calc(100vh-1.5rem)]">
        <div
            class="min-w-[175px] max-w-[50%] mt-1"
            style="width: {resizeWidth}px;"
        >
            <Sidebar />
        </div>
        <div
            role="separator"
            aria-orientation="vertical"
            class="w-2 cursor-col-resize"
            use:onDrag={{ orientation: "vertical" }}
            on:drag={({ detail: delta }) => handleDrag(delta)}
            on:dragEnd={() => {
                width = resizeWidth;
            }}
        ></div>
        <div class="flex flex-col flex-grow mr-2 mb-2">
            <div class="flex-none h-16 rounded-md mt-1"><Header recipient={"Some Person"} /></div>
            <div class="bg-gray-700 flex-grow rounded-md mt-2"></div>
            <div class="flex-none h-24 rounded-md mt-2"><Footer recipient={"Some Person"} /></div>
        </div>
    </div>
</main>
