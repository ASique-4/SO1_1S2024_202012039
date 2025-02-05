#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>
//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>	
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Hoja de Trabajo 1, Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Angel Sique");

static int escribir_archivo(struct seq_file *archivo, void *v)
{
    seq_printf(archivo, "{\"data\":\"");
    seq_printf(archivo, "*********************************************\n");
    seq_printf(archivo, "**    LABORATORIO SISTEMAS OPERATIVOS 1    **\n");
    seq_printf(archivo, "**            HOJA DE TRABAJO 1            **\n");
    seq_printf(archivo, "**               ANGEL SIQUE               **\n");
    seq_printf(archivo, "*********************************************\n");
    seq_printf(archivo, "\"}\n");
    return 0;
}


//Funcion que se ejecuta cuando se le hace un cat al modulo.
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

// Si el su Kernel es 5.6 o mayor
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

static int _insert(void)
{
    proc_create("modulo_HT1", 0, NULL, &operaciones);
    printk(KERN_INFO "Mensaje al insertar el modulo\n");
    return 0;
}

static void _remove(void)
{
    remove_proc_entry("modulo_HT1", NULL);
    printk(KERN_INFO "Mensaje al remover el modulo\n");
}

module_init(_insert);
module_exit(_remove);