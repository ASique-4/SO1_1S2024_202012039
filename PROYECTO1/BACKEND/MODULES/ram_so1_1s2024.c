#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sysinfo.h>
#include <linux/seq_file.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel Sique");
MODULE_DESCRIPTION("Informacion de la RAM");
MODULE_VERSION("1.0");

struct sysinfo inf;

static int write_to_proc(struct seq_file *file_proc, void *v)
{
    unsigned long total, used, notused;
    unsigned long porc;
    si_meminfo(&inf);

    total = inf.totalram * inf.mem_unit;
    used = (inf.totalram - inf.freeram) * inf.mem_unit;
    porc = (used * 100) / total;
    notused = total - used;

    seq_printf(file_proc, "{\"totalRam\":%lu, \"memoriaEnUso\":%lu, \"porcentaje\":%lu, \"libre\":%lu }", total, used, porc, notused);
    return 0;
}

static int open_proc(struct inode *inode, struct file *file)
{
    return single_open(file, write_to_proc, NULL);
}

static const struct file_operations proc_fops = {
    .open = open_proc,
    .read = seq_read,
    .llseek = seq_lseek,
    .release = single_release,
};

static int __init ram_module_init(void)
{
    proc_create("ram_so1_1s2024", 0, NULL, &proc_fops);
    printk(KERN_INFO "Módulo RAM montado\n");
    return 0;
}

static void __exit ram_module_exit(void)
{
    remove_proc_entry("ram_so1_1s2024", NULL);
    printk(KERN_INFO "Módulo RAM eliminado\n");
}

module_init(ram_module_init);
module_exit(ram_module_exit);
