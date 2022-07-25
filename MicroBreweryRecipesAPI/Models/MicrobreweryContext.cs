using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;

namespace MicroBreweryRecipesAPI.Models
{
    public class MicroBreweryContext : DbContext
    {
        public DbSet<Recipe> Recipes {get;set;}
        public DbSet<Ingredient> Ingredients{get;set;}
        public DbSet<IngredientCategory> IngredientCategories {get;set;}
        public DbSet<RecipeCategory> RecipeCategories{get;set;}
        public DbSet<RecipeIngredientList> RecipeIngredientLists {get;set;}

        public MicroBreweryContext(DbContextOptions<MicroBreweryContext> options) : base(options)
        {

        }
        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
            => optionsBuilder.UseNpgsql("Host=127.0.0.1;Database=microbrewery;Username=postgres;Password=kernel21");
    }
}
