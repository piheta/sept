<script>
    // @ts-nocheck
    import Sidebar from "/src/components/Sidebar/Sidebar.svelte";
    import Footer from "/src/components/Footer.svelte";
    import { onDrag } from "/src/components/Sidebar/SidebarResizer";
    import Header from "/src/components/Header.svelte";
    import Messages from "/src/components/Chatbox.svelte";
    import Chatbox from "/src/components/Chatbox.svelte";
    
    let width = 180;
    let resizeWidth = width;
    const minWidth = 180;
    const maxWidth = () => window.innerWidth / 2;

    $: console.log(resizeWidth);

    function handleDrag(delta) {
        const newWidth = width + delta;
        resizeWidth = Math.max(minWidth, Math.min(maxWidth(), newWidth));
    }
</script>

<main class="h-full select-none">
    <div class="wails-drag w-full h-6 absolute top-0 left-0"></div>
    <div class="flex w-full h-[calc(100vh)]">
        <div
            class="min-w-[180px] max-w-[50%] mt-5"
            style="width: {resizeWidth}px;"
        >
            <Sidebar />
        </div>
        <div
            role="separator"
            aria-orientation="vertical"
            class="w-2 min-w-2 cursor-col-resize"
            use:onDrag={{ orientation: "vertical" }}
            on:drag={({ detail: delta }) => handleDrag(delta)}
            on:dragEnd={() => {
                width = resizeWidth;
            }}
        ></div>
        <div class="flex flex-col flex-grow mr-2 mb-2 mt-2.5">

                <Chatbox />
            <div
                role="separator"
                class="h-2 min-h-2"
            ></div>
            <!-- <div class="flex-none rounded-md" style="height: {resizeFooterHeight}px;"> -->
                <Footer recipient={"Some Person"} height={120} />
            <!-- </div> -->
        </div>
    </div>
</main>
