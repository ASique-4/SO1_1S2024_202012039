#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sysinfo.h>

static int __init ram_module_init(void) {
    struct sysinfo info;

    si_meminfo(&info);

    printk(KERN_INFO "RAM Total: %lu KB\n", (info.totalram * (unsigned long)info.mem_unit) / 1024);
    printk(KERN_INFO "RAM Libre: %lu KB\n", (info.freeram * (unsigned long)info.mem_unit) / 1024);

    return 0;
}

static void __exit ram_module_exit(void) {
    printk(KERN_INFO "Exiting RAM module\n");
}

module_init(ram_module_init);
module_exit(ram_module_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel Sique");
MODULE_DESCRIPTION("Módulo de kernel para mostrar información sobre la RAM");
