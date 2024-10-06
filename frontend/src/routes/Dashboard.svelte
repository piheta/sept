<script>
    // @ts-nocheck
    import Sidebar from "/src/components/Sidebar/Sidebar.svelte";
    import Footer from "/src/components/Footer.svelte";
    import { onDrag } from "/src/components/Sidebar/SidebarResizer";
    import Header from "/src/components/Header.svelte";
    import Chatbox from "/src/components/Chatbox.svelte";
    
    let width = 180;
    let resizeWidth = width;
    const snapWidth = 70; 
    const snapThreshold = 110;
    const minWidth = 180;
    const maxWidth = () => window.innerWidth / 2;
    
    $: console.log(resizeWidth);
    
    function handleDrag(delta) {
        const newWidth = width + delta;
    
        // If the user drags below the snapping threshold (140px), snap to 60px
        if (newWidth < snapThreshold) {
            resizeWidth = snapWidth;
        } 
        // If the new width is between 140px and 180px, hold at 180px visually
        else if (newWidth < minWidth) {
            resizeWidth = minWidth;
        } 
        // Otherwise, allow normal resizing above 180px
        else {
            resizeWidth = Math.min(maxWidth(), newWidth);
        }
    }
    </script>
    
    <main class="h-full select-none">
        <div class="wails-drag w-full h-6 absolute top-0 left-0"></div>
        <div class="flex w-full h-[calc(100vh)]">
            <div
                class="min-w-[60px] max-w-[50%] mt-5"
                style="width: {resizeWidth}px;"
            >
                <Sidebar small={resizeWidth===70} />
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
                <div role="separator" class="h-2 min-h-2"></div>
                <Footer recipient={"Some Person"} height={120} />
            </div>
        </div>
    </main>
    