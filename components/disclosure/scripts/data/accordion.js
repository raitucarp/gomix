function accordion(initData = {}) {
  return {
    openItems: initData.openItems ?? [],
    multiple: initData.multiple ?? false,
    collapsible: initData.collapsible ?? true,

    // init() {
    //   if (this.openItems.length === 0 && !this.collapsible) {
    //     this.openItems = [0];
    //   }
    //   this.$watch("multiple", (value) => {
    //     if (!value && this.openItems.length > 1) {
    //       this.openItems = [this.openItems[0]];
    //     }
    //   });
    //   this.$watch("collapsible", (value) => {
    //     if (!value && this.openItems.length === 0) {
    //       this.openItems = [0];
    //     }
    //   });
    // },

    isOpen(item) {
      return this.openItems.includes(item);
    },

    toggle(item) {
      if (this.isOpen(item)) {
        if (this.collapsible) {
          this.openItems = this.openItems.filter((i) => i !== item);
        }
      } else {
        if (this.multiple) {
          this.openItems.push(item);
        } else {
          this.openItems = [item];
        }
      }
    },
  };
}
