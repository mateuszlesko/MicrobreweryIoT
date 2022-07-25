using System.Collections.Generic;

namespace MicroBreweryRecipesAPI.Models
{
    public class RecipeCategory
    {

        //entity properties
        public int Id {get;set;}
        public string RecipeCategoryName {get;set;}

        //relationships
        public List<Recipe> Recipes {get;set;}
    }
}